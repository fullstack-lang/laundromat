// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { WasherDB } from './washer-db';

// insertion point for imports
import { MachineDB } from './machine-db'

@Injectable({
  providedIn: 'root'
})
export class WasherService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  WasherServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private washersUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.washersUrl = origin + '/api/github.com/fullstack-lang/laundromat/go/v1/washers';
  }

  /** GET washers from the server */
  getWashers(): Observable<WasherDB[]> {
    return this.http.get<WasherDB[]>(this.washersUrl)
      .pipe(
        tap(_ => this.log('fetched washers')),
        catchError(this.handleError<WasherDB[]>('getWashers', []))
      );
  }

  /** GET washer by id. Will 404 if id not found */
  getWasher(id: number): Observable<WasherDB> {
    const url = `${this.washersUrl}/${id}`;
    return this.http.get<WasherDB>(url).pipe(
      tap(_ => this.log(`fetched washer id=${id}`)),
      catchError(this.handleError<WasherDB>(`getWasher id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new washer to the server */
  postWasher(washerdb: WasherDB): Observable<WasherDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    washerdb.Machine = new MachineDB

    return this.http.post<WasherDB>(this.washersUrl, washerdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`posted washerdb id=${washerdb.ID}`)
      }),
      catchError(this.handleError<WasherDB>('postWasher'))
    );
  }

  /** DELETE: delete the washerdb from the server */
  deleteWasher(washerdb: WasherDB | number): Observable<WasherDB> {
    const id = typeof washerdb === 'number' ? washerdb : washerdb.ID;
    const url = `${this.washersUrl}/${id}`;

    return this.http.delete<WasherDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted washerdb id=${id}`)),
      catchError(this.handleError<WasherDB>('deleteWasher'))
    );
  }

  /** PUT: update the washerdb on the server */
  updateWasher(washerdb: WasherDB): Observable<WasherDB> {
    const id = typeof washerdb === 'number' ? washerdb : washerdb.ID;
    const url = `${this.washersUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    washerdb.Machine = new MachineDB

    return this.http.put<WasherDB>(url, washerdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated washerdb id=${washerdb.ID}`)
      }),
      catchError(this.handleError<WasherDB>('updateWasher'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
