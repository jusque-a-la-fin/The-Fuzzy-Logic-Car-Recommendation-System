export function doRedirection(href, hrefError) {
    fetch(href)
    .then(response => {
      if (response.ok) {
        window.location.href = href;
      } else {
        window.location.href = hrefError + "?code=" + response.status;
        throw new Error('HTTP Response Status Code: ' + response.status);
      }})
      .catch(error => console.error(error));
  }
  
export function processError(href) {
    fetch(href, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ message: "sessionID wasn't set in Session Storage" })
    })
    .then(response => {
      if (!response.ok) {
        window.location.href = href + "?code=" + response.status;
        throw new Error('HTTP Response Status Code: ' + response.status);
      }
    })
    .catch(error => console.error(error));
  }