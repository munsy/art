import { Component } from '@angular/core';

import { NgForm } from '@angular/forms'; 
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { ArtifactoryService } from './artifactory.service';
import { ArtifactList, ArtifactoryRequest } from './artifactory.model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public files: ArtifactList;
  public results: boolean = false;
  public searching: boolean = false;
  public error: boolean = false;

  constructor(private service: ArtifactoryService) { }

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

  	this.service.getList(request).subscribe(response => {
      this.files = response as ArtifactList;
      this.results = true;
      console.log(response);
  		console.log(this.files);
  	}, error => {
      this.error = true;
      this.results = false;
      this.searching = false;
      console.log(error);
    });
  }
}
