let loc = (): string => { 
  var winloc = window.location.href;
  return winloc.endsWith('/') ? winloc.slice(0, -1) : winloc;
}

export const environment = {
  production: true,
  apiURL: loc()
};
