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

  constructor(private service: ArtifactoryService) { }

  getResults(form: NgForm) {
  	console.log(form.value);
  	
  	let request = new ArtifactoryRequest();
  	request.url = form.value.url;
  	request.repo = form.value.repo;
  	request.username = form.value.username;
  	request.password = form.value.password;

  	this.service.getList(request).subscribe(response => {
  		this.results = true;
  		console.log(response);
  	});
  }
}
