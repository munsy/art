import { Component } from '@angular/core';

import { NgForm } from '@angular/forms'; 
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { ArtifactoryService } from './artifactory.service';
import { ArtifactList, ArtifactResult, ArtifactoryRequest } from './artifactory.model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public files: Array<ArtifactResult>;
  public results: boolean = false;
  public searching: boolean = false;
  public error: boolean = false;
  public filter: boolean = false;

  constructor(private service: ArtifactoryService) { }

  toggle(e: any) {
    this.filter = !this.filter;
  }

  downloadCount(result: ArtifactResult): number {
    let count = 0;
    for(var i = 0; i < result.stats.length; i++) {
        count += result.stats[i].downloads;
    }
    return count;
  }

  mostDownloaded(): Array<ArtifactResult> {
    let first = 0;
    let second = 0;

    for(var i = 0; i < this.files.length; i++) {
      if(this.downloadCount(this.files[i]) > first) {
        second = first;
        first = this.downloadCount(this.files[i]);
      }
    }

    let list = new Array<ArtifactResult>();
    for(var i = 0; i < this.files.length; i++) {
      if(this.downloadCount(this.files[i]) == first) {
        list.push(this.files[i]);
      }
    }
    for(var i = 0; i < this.files.length; i++) {
      if(this.downloadCount(this.files[i]) == second) {
        list.push(this.files[i]);
      }
    }
    
    return list;
  }

  reset() {
    this.results = false;
    this.searching = false;
    this.error = false;
  }

  getResults(form: NgForm) {
    this.searching = true;
    this.results = false;
  	let request = new ArtifactoryRequest();
  	request.url = form.value.url;
  	request.repo = form.value.repo;
  	request.username = form.value.username;
  	request.password = form.value.password;

    console.log(request);

  	this.service.getList(request).subscribe(response => {
      let list = response as ArtifactList;
      this.files = list.results;
      console.log(this.files);
      this.results = true;
      this.searching = false;
  	}, error => {
      this.error = true;
      this.results = false;
      this.searching = false;
    });
  }

  sort() {
    console.log("not implemented yet");
  }
}
