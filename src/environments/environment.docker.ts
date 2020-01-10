export const environment = {
  production: true,
  apiURL: (): string => { 
  	var loc = window.location.href;
  	return loc.endsWith('/') ? loc.slice(0, -1) : loc;
  }
};
