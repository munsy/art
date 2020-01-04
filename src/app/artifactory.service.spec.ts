import { TestBed } from '@angular/core/testing';

import { ArtifactoryService } from './artifactory.service';

describe('ArtifactoryService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: ArtifactoryService = TestBed.get(ArtifactoryService);
    expect(service).toBeTruthy();
  });
});
