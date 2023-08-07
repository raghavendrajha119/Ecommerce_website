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
      })
      .catch(error => {
        console.error('Error fetching data: ', error);
      });
  }
fetchData();  