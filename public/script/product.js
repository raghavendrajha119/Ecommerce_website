function seedetails(productId){
  window.location.href = `/product.html?id=${productId}`;
}
function fetchData() {
    const urlParams = new URLSearchParams(window.location.search);
    const productId = urlParams.get('id');
    fetch(`/products?id=${productId}`)
      .then(response => response.json())
      .then(product => {
        const productContainer = document.querySelector('.container');
        productContainer.innerHTML = '';
  
        const card = document.createElement('div');
        card.classList.add('card');
        const prodImg = document.createElement('div');
        prodImg.classList.add('img');
        prodImg.innerHTML = `
          <img src="./img/${product.Title}.jpg" alt="${product.Title}">
        `;
        const prodDiv = document.createElement('div');
        prodDiv.classList.add('details');
        prodDiv.innerHTML = `
          <h2>${product.Title}</h2>
          <p class="cat">${product.Category}</p>
          <h5>Price: INR ${product.Price}</h5>
          <p class="desc">${product.Description}</p>
          <div class="buttons">
            <button class="addtocart" type="submit">Add to Bag</button>
            <button class="Buynow" type="submit">Buy Now</button>
          </div>
        `;
        card.appendChild(prodImg);
        card.appendChild(prodDiv);
        productContainer.appendChild(card);
        // Fetch similar products based on the current product's details
      fetch(`/similar-products?title=${product.Title}&description=${product.Description}&category=${product.Category}`)
      .then(response => response.json())
      .then(similarProducts => {
        displaySimilarItemsSuggestions(similarProducts);
      });
      })
      .catch(error => {
        console.error('Error fetching data: ', error);
      });
  }
fetchData();

// Function to display similar items suggestions
function displaySimilarItemsSuggestions(suggestions) {
  const similarItemsContainer = document.querySelector('.similar-items-container');
  similarItemsContainer.innerHTML = '';
  console.log(suggestions)
  suggestions.forEach(suggestion => {
      const suggestionDiv = document.createElement('div');
      suggestionDiv.classList.add('similar-item');
      suggestionDiv.innerHTML = `
        <div class="card">
          <img src="./img/${suggestion.Title}.jpg" alt="${suggestion.Title}">
          <p>${suggestion.Title}</p>
          <button class="seedetails" type="submit" onclick="seedetails('${suggestion.ID}')">Details</button>
        </div>
      `;
      similarItemsContainer.appendChild(suggestionDiv);
  });
}