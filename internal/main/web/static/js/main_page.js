import { doRedirection, processError } from '../shared-js/utils.js';

const preloader = document.getElementById("preloader");
preloader.style.display = "none";

 const brandSelect = document.getElementById('make');
 const modelSelect = document.getElementById('model');

 const initialModelOptions = modelSelect.innerHTML;


 brandSelect.addEventListener('change', function () {
   const selectedBrandId = brandSelect.options[brandSelect.selectedIndex].id;
   modelSelect.innerHTML = initialModelOptions;

   if (!selectedBrandId) {
	 modelSelect.disabled = true;
   } else {
	 modelSelect.disabled = false;
	 Array.from(modelSelect.options).forEach(function (option) {
	   const brandId = option.getAttribute('id'); 
	   if (brandId === selectedBrandId) {
		 option.style.display = 'block'; 
	   } else {
		 option.style.display = 'none'; 
	   }
	 });
   }
 });



const lowPriceInput = document.getElementById('low_price_limit');
const highPriceInput = document.getElementById('high_price_limit');

lowPriceInput.addEventListener('blur', () => {
const value = lowPriceInput.value.trim();

if (isNaN(value) || value < 0 || value[0] === '0') {
lowPriceInput.value = ''; 
lowPriceInput.classList.add('is-invalid');
document.getElementsByClassName('error low_price')[0].innerText = 'Введите число больше 0';
} else {
lowPriceInput.classList.remove('is-invalid');
document.getElementsByClassName('error low_price')[0].innerText = '';
}
});

highPriceInput.addEventListener('blur', () => {
const value = highPriceInput.value.trim();
if (isNaN(value) || value < 0 || value[0] === '0') {
highPriceInput.value = ''; 
highPriceInput.classList.add('is-invalid');
document.getElementsByClassName('error high_price')[0].innerText = 'Введите число больше 0';
} else {
highPriceInput.classList.remove('is-invalid');
document.getElementsByClassName('error high_price')[0].innerText = '';
}
});


const div = document.getElementById("fuzzy_algorithm");
div.addEventListener("click", getSelection);


function getSelection() {
    const sessionID = sessionStorage.getItem('sessionID');
    if (sessionID) {
        window.location.href = "http://localhost:8082/selection/priorities/guest/"+sessionID;
    } else {
      const href = 'http://localhost:8080/main/error';
      processError(href);
    }
}

const button = document.getElementById('searchButton');
button.addEventListener('click', doSearching);

function doSearching() {
    const sessionID = sessionStorage.getItem('sessionID');
    if (sessionID) {
      hideAllAndShowLoading();
        var isKnown=false;
        if (localStorage.getItem('isknown')==='true') {
            isKnown = true;
        } else {
          localStorage.setItem('isknown', 'true');
        }

        const form = document.getElementById("usual_search")
        const formData = new FormData(form);
        let formDataObject = {};
        formData.forEach(function(value, key){
        formDataObject[key] = value;
        });

        const data = { sessionID: sessionID, form: formDataObject, isKnown: isKnown};
        const url = 'http://localhost:8081/search';
        fetch(url, {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(response => {
            if (response.ok) {
              const href = "http://localhost:8081/search/guest/"+sessionID+"?survey="+true+"&isknown="+isKnown;
           
              const hrefError = "http://localhost:8081/search/error";
              doRedirection(href, hrefError);
            } else {
              window.location.href = "http://localhost:8081/search/error/code/"+response.status;
              throw new Error('HTTP Response Status Code: ' + response.status);  
            }
        })
        .catch(error => console.error(error));
    } else {
      const href = 'http://localhost:8081/search/error';
      processError(href);
    }
}

function hideAllAndShowLoading() {
  const form = document.getElementById("usual_search")
	form.style.display = "none";
	const fuzzy_algorithm = document.getElementById("fuzzy_algorithm")
	fuzzy_algorithm.style.display = "none"
	preloader.style.display = "block";
}