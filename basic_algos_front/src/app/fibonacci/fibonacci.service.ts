import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { Observable, of, throwError } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';



@Injectable({
  providedIn: 'root'
})
export class FibonacciService {

  private utilUrl = 'http://localhost:8080/fibonacci/';  // URL to web api

  constructor(private http: HttpClient,) { }

  get(): Observable<string> {
    return this.http.get(this.utilUrl, { responseType: 'text' })
      .pipe(catchError(this.handleError));
  }

  post(n: number): Observable<string> {
    var formData = new FormData()
    formData.append("n", n.toString());

    return this.http.post(this.utilUrl, formData, { responseType: 'text' } )
      .pipe(catchError(this.handleError));
  }

  /**
   * Log the message using some system
   * @param message 
   */
  private log(message: string) {
    console.log(`${message}`);
  }

  private handleError(httpError: HttpErrorResponse) {
    if (httpError.error instanceof ErrorEvent) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', httpError.error.message);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong.
      console.error(
        `Backend returned code ${httpError.status}, ` +
        `body was: ${httpError.error}`);
    }
    // Return an observable with a user-facing error message.
    return throwError('Something bad happened; please try again later.');
  }
}
