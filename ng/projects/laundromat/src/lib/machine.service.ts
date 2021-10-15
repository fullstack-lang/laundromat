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

import { MachineDB } from './machine-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class MachineService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  MachineServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private machinesUrl: string

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
    this.machinesUrl = origin + '/api/github.com/fullstack-lang/laundromat/go/v1/machines';
  }

  /** GET machines from the server */
  getMachines(): Observable<MachineDB[]> {
    return this.http.get<MachineDB[]>(this.machinesUrl)
      .pipe(
        tap(_ => this.log('fetched machines')),
        catchError(this.handleError<MachineDB[]>('getMachines', []))
      );
  }

  /** GET machine by id. Will 404 if id not found */
  getMachine(id: number): Observable<MachineDB> {
    const url = `${this.machinesUrl}/${id}`;
    return this.http.get<MachineDB>(url).pipe(
      tap(_ => this.log(`fetched machine id=${id}`)),
      catchError(this.handleError<MachineDB>(`getMachine id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new machine to the server */
  postMachine(machinedb: MachineDB): Observable<MachineDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    return this.http.post<MachineDB>(this.machinesUrl, machinedb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`posted machinedb id=${machinedb.ID}`)
      }),
      catchError(this.handleError<MachineDB>('postMachine'))
    );
  }

  /** DELETE: delete the machinedb from the server */
  deleteMachine(machinedb: MachineDB | number): Observable<MachineDB> {
    const id = typeof machinedb === 'number' ? machinedb : machinedb.ID;
    const url = `${this.machinesUrl}/${id}`;

    return this.http.delete<MachineDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted machinedb id=${id}`)),
      catchError(this.handleError<MachineDB>('deleteMachine'))
    );
  }

  /** PUT: update the machinedb on the server */
  updateMachine(machinedb: MachineDB): Observable<MachineDB> {
    const id = typeof machinedb === 'number' ? machinedb : machinedb.ID;
    const url = `${this.machinesUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    return this.http.put<MachineDB>(url, machinedb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated machinedb id=${machinedb.ID}`)
      }),
      catchError(this.handleError<MachineDB>('updateMachine'))
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
