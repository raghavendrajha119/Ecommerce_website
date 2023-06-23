const rangeInput = document.getElementById('myRange');
const minusButton = document.querySelector('.minus-button');
const plusButton = document.querySelector('.plus-button');
const rangeValue = document.getElementById('range-value');

minusButton.addEventListener('click', () => {
  rangeInput.value = parseInt(rangeInput.value) - 1;
  rangeValue.textContent = rangeInput.value;
});

plusButton.addEventListener('click', () => {
  rangeInput.value = parseInt(rangeInput.value) + 1;
  rangeValue.textContent = rangeInput.value;
});
