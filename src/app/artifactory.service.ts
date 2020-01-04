import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, interval } from 'rxjs';
import { timeout } from 'rxjs/operators';

import { ArtifactoryRequest, ArtifactList } from './artifactory.model';

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
  constructor(private http: HttpClient) { }

  getList(request: ArtifactoryRequest): Observable<ArtifactList> {
    var endpoint = `http://${request.url}/artifactory/api/search/aql`;
    httpOptions.headers = httpOptions.headers.set('Authorization', btoa(`${request.username}:${request.password}`));
    var query = `items.find({"repo":{"$eq":"${request.repo}"}}).include("stat")`;
    return this.http.post<ArtifactList>(endpoint, query, httpOptions).pipe(timeout(5000));
  }
}
