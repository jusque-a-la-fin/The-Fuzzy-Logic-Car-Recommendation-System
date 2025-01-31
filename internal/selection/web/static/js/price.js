function isNumber(value) {
  return !isNaN(parseFloat(value)) && isFinite(value);
}

function validateInput(inputElement, errorElement) {
  let value = inputElement.value;
  if (value.startsWith('0')) {
    value = '';
  }

  if (isNumber(value) && value >= 0) {
      inputElement.classList.remove('error');
      errorElement.textContent = '';
      inputElement.placeholder = 'Самая низкая цена';
  } else {
      inputElement.classList.add('error');
      errorElement.textContent = 'Введите число больше нуля';
      inputElement.value = ''; 
      inputElement.placeholder = ''; 
  }
}

const lstPriceInput = document.getElementById('lst_price');
const hstPriceInput = document.getElementById('hst_price');
const lstPriceError = document.getElementById('lst_price_error');
const hstPriceError = document.getElementById('hst_price_error');

lstPriceInput.addEventListener('input', function() {
  validateInput(lstPriceInput, lstPriceError);
});

hstPriceInput.addEventListener('input', function() {
  validateInput(hstPriceInput, hstPriceError);
});

function sendRequest() {
  const data = { minPrice: document.getElementById('lst_price').value, maxPrice: document.getElementById('hst_price').value};
  sessionStorage.setItem('prices', JSON.stringify(data));
  const url = window.location.href;
  const sessionID = url.match(/\/selection\/price\/guest\/([^/?]+)/)[1];
  if (sessionID) {
    window.location.href = "http://localhost:8082/selection/manufacturers/guest/"+sessionID;
  } else {
    fetch('http://localhost:8082/selection/error', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ message: "sessionID wasn't set in Session Storage" })
    })
    .then(response => {
      if (!response.ok) {
      window.location.href = "http://localhost:8082/selection/error?code="+response.status;
      throw new Error('HTTP Response Status Code: ' + response.status);
      }
    })
    .catch(error => console.error(error));
  }
}