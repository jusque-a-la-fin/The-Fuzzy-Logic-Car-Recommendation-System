const divElements = document.querySelectorAll(".car_page");

divElements.forEach(function(divElement) {
    divElement.addEventListener('click', getCarPage);
});

function getCarPage(event) {
    const divElement = event.currentTarget;
    const carID = divElement.getAttribute("carID");
    const errorLink = divElement.getAttribute("errorLink");
    const href = constructURLOfCarPage(window.location.href, carID)
   
    fetch(href)
    .then(response => {
        if (response.ok) {
            window.location.href = href;
        } else {
            window.location.href = errorLink + response.status;
            throw new Error('HTTP Response Status Code: ' + response.status);
        }
    })
    .catch(error => console.error(error));
}

function constructURLOfCarPage(url, cardID) {
    const urlObj = new URL(url);
    const pathParts = urlObj.pathname.split('/');
    const guestIndex = pathParts.indexOf('guest');

    if (guestIndex !== -1 && guestIndex + 1 < pathParts.length) {
        pathParts.splice(guestIndex + 2, 0, 'carID', cardID);
    }

    urlObj.pathname = pathParts.join('/');
    return urlObj.toString();
}