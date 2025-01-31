import { doRedirection, processError } from '../shared-js/utils.js';

const preloader = document.getElementById('preloader');
preloader.style.display = "none";

const buttonOK = document.getElementById('button_ok');
buttonOK.addEventListener('click', sendRequest);

var btns = document.getElementsByClassName("answer");
const btnsArray = Array.from(btns);

const countires = ['Германия', 'Япония', 'Россия', 'Китай', 'США', 'Великобритания', 'Франция', 'Южная_Корея', 'Другие'];
var choices = [];
var clicks = [];

btnsArray.forEach((btn, index) => {
  btn.addEventListener("click", changeStyleWhenClick);
  btn.addEventListener('click', function() {
    changeChoice(countires[index], index);
  });
});

function changeStyleWhenClick() {
    if (this.classList.length <= 2) {
    this.classList.add("clicked");

  } else {
    this.classList.remove("clicked");
  }
}

function changeChoice(choice, index) {
    if (clicks[index] == undefined) {
      clicks[index] = true;
      choices.push(choice);
    }
    else {
      clicks[index] = undefined;
      let indexOfChoicesElement = choices.indexOf(choice);
      choices.splice(indexOfChoicesElement, 1);
    }
}

function sendRequest() {
  const url = window.location.href;
  const sessionID = url.match(/\/selection\/manufacturers\/guest\/([^/?]+)/)[1];
  if (sessionID) {
    goToTheDestination();
  } else {
    const href = 'http://localhost:8082/selection/error';
    processError(href);
  }
}

function goToTheDestination() {
    let url = window.location.href;
    const sessionID = url.match(/\/selection\/manufacturers\/guest\/([^/?]+)/)[1];
    const priorities = JSON.parse(sessionStorage.getItem('priorities')).priorities;
    const minPrice = JSON.parse(sessionStorage.getItem('prices')).minPrice;
    const maxPrice = JSON.parse(sessionStorage.getItem('prices')).maxPrice;
    const data = {priorities: priorities, minPrice: minPrice, maxPrice: maxPrice, manufacturers: choices, sessionID: sessionID}
    hideAllAndShowLoading();
      url = 'http://localhost:8082/selection';
      fetch(url, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
      })
      .then(response => {
          if (response.ok) {
            const href = "http://localhost:8082/selection/guest/"+sessionID;
            const hrefError = "http://localhost:8082/selection/error";
            doRedirection(href, hrefError);
          } else {
            window.location.href = "http://localhost:8082/selection/error?code="+response.status;
            throw new Error('HTTP Response Status Code: ' + response.status);
          }
      })
      .catch(error => console.error(error));
}

function hideAllAndShowLoading() {
    var elements = document.getElementsByTagName('*');
    for (var i = 0; i < elements.length; i++) {
      elements[i].style.display = 'none';
    }
    preloader.style.display = "block";
}