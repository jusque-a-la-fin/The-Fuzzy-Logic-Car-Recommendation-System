import { doRedirection } from '../shared-js/utils.js';

const urlParams = new URLSearchParams(window.location.search);

if(urlParams.get('isknown')===true) {
    document.cookie = `userID=${localStorage.getItem('userID')}; SameSite=Strict;`;
}

function setUserID() {
if (!localStorage.getItem('userID')) {
        const cookieUserID = document.cookie.replace(/(?:(?:^|.*;\s*)userID\s*\=\s*([^;]*).*$)|^.*$/, "$1");
        if (cookieUserID) {
            localStorage.setItem('userID', cookieUserID);
        }
}
}


if (urlParams.get('survey')) {
    const questionID = document.getElementById("question").getAttribute("questionID");
    sessionStorage.setItem('questionID', questionID);
}

const form = document.querySelector('form');
if (form !== null) {
form.addEventListener('submit', function(event) {
    event.preventDefault();

    const formData = new FormData(form);
    const answer = formData.get('radio');
    const questionID = sessionStorage.getItem('questionID');
    
    const data = {
        questionID: questionID,
        answer: answer,
    };
    
    fetch('http://localhost:8081/search/answer', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .then(response => {
        if (response.ok) {
            setUserID();

            const url = window.location.href;
            const sessionID = url.match(/\/search\/guest\/([^/?]+)/)[1];
          
            const href = "http://localhost:8081/search/guest/"+sessionID;
          
            const hrefError = "http://localhost:8081/search/error";
            doRedirection(href, hrefError);
           
    } else {
        window.location.href = "http://localhost:8081/search/error?code="+response.status;
        throw new Error('HTTP Response Status Code: ' + response.status);  
    }})
    .catch(error => console.error(error));
});
}