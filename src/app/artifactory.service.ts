import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, interval } from 'rxjs';
import { timeout } from 'rxjs/operators';

import { ArtifactoryRequest, ArtifactList } from './artifactory.model';

import { environment } from '../environments/environment';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json',
    'Access-Control-Allow-Origin': '*',
  })
};

@Injectable({
  providedIn: 'root'
})
export class ArtifactoryService {
  private url: string = environment.apiURL;
  constructor(private http: HttpClient) { }

  getList(request: ArtifactoryRequest): Observable<ArtifactList> {
    return this.http.post<ArtifactList>(`${this.url}/api/v1/artifactory`, request).pipe(timeout(5000));
  }
}
