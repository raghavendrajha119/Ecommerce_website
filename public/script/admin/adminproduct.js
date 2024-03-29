function redirectToEditPage(productID) {
  window.location.href = `./editproduct.html?id=${productID}`;
} 
function fetchData() {
    fetch(`/admin/products`)
    .then(response => response.json())
    .then(products => {
        const productContainer = document.querySelector('.container');
        products.forEach(product => {
        const card = document.createElement('div');
        card.classList.add('card');
        const prodImg = document.createElement('div');
        prodImg.classList.add('img');
        prodImg.innerHTML = `
          <img src="../img/${product.Title}.jpg" alt="${product.Title}">
        `;
        const prodDiv = document.createElement('div');
        prodDiv.classList.add('details');
        prodDiv.innerHTML = `
          <h2>${product.Title}</h2>
          <p class="cat">${product.Category}</p>
          <h5>Price: INR ${product.Price}</h5>
          <p>Quantity: ${product.Quantity}</p>
          <p class="desc">${product.Description}</p>
          <div class="buttons">
            <button class="Edit" type="submit" onclick="redirectToEditPage('${product.ID}')">Edit</button>  
          </div>
        `;
        card.appendChild(prodImg);
        card.appendChild(prodDiv);
        productContainer.appendChild(card);
        });
      })
      .catch(error => {
        console.error('Error fetching data: ', error);
      });
  }
fetchData();
 