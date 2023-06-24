// Update the URL of the fetch request based on your server configuration
fetch('/cart-products')
  .then(response => response.json())
  .then(data => {
    const productIDs = data.productIDs;
    // count for the repeated products
    const productQuantity = {};
    productIDs.forEach(productID => {
      productQuantity[productID] = productQuantity[productID] + 1 || 1;
    });
    const uniqueProductIDs = [...new Set(productIDs)];
    // Use the product IDs to fetch the product details from the FakeStore API
    Promise.all(uniqueProductIDs.map(productID =>
      fetch(`https://fakestoreapi.com/products/${productID}`)
        .then(response => response.json())
    ))
      .then(products => {
        // Display the products in the cart page
        const cartItemsContainer = document.querySelector('.cart-items');
        let totalPrice = 0;
        products.forEach(product => {
          const productHTML = `
            <div class="cart-item">
              <img src="${product.image}" alt="${product.title}">
              <h3 class="name">${product.title}</h3>
              <p class="quantity">Quantity: ${productQuantity[product.id]}</p>
              <p class="price">$${product.price}</p>
             
            </div>
          `;
          cartItemsContainer.innerHTML += productHTML;
          totalPrice += product.price;
        });
        const totalAmountElement = document.getElementById('totalAmount');
        totalAmountElement.textContent = `$${totalPrice.toFixed(2)}`;
      })
      .catch(error => console.error(error));
  })
  .catch(error => console.error(error));
