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

import { AgentDB } from './agent-db';

// insertion point for imports
import { EngineDB } from './engine-db'

@Injectable({
  providedIn: 'root'
})
export class AgentService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  AgentServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private agentsUrl: string

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
    this.agentsUrl = origin + '/api/github.com/fullstack-lang/gongsim/go/v1/agents';
  }

  /** GET agents from the server */
  getAgents(): Observable<AgentDB[]> {
    return this.http.get<AgentDB[]>(this.agentsUrl)
      .pipe(
        tap(_ => this.log('fetched agents')),
        catchError(this.handleError<AgentDB[]>('getAgents', []))
      );
  }

  /** GET agent by id. Will 404 if id not found */
  getAgent(id: number): Observable<AgentDB> {
    const url = `${this.agentsUrl}/${id}`;
    return this.http.get<AgentDB>(url).pipe(
      tap(_ => this.log(`fetched agent id=${id}`)),
      catchError(this.handleError<AgentDB>(`getAgent id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new agent to the server */
  postAgent(agentdb: AgentDB): Observable<AgentDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    agentdb.Engine = new EngineDB

    return this.http.post<AgentDB>(this.agentsUrl, agentdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`posted agentdb id=${agentdb.ID}`)
      }),
      catchError(this.handleError<AgentDB>('postAgent'))
    );
  }

  /** DELETE: delete the agentdb from the server */
  deleteAgent(agentdb: AgentDB | number): Observable<AgentDB> {
    const id = typeof agentdb === 'number' ? agentdb : agentdb.ID;
    const url = `${this.agentsUrl}/${id}`;

    return this.http.delete<AgentDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted agentdb id=${id}`)),
      catchError(this.handleError<AgentDB>('deleteAgent'))
    );
  }

  /** PUT: update the agentdb on the server */
  updateAgent(agentdb: AgentDB): Observable<AgentDB> {
    const id = typeof agentdb === 'number' ? agentdb : agentdb.ID;
    const url = `${this.agentsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    agentdb.Engine = new EngineDB

    return this.http.put<AgentDB>(url, agentdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated agentdb id=${agentdb.ID}`)
      }),
      catchError(this.handleError<AgentDB>('updateAgent'))
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
