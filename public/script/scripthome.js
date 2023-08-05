function fetchproducts() {
  fetch('/')
    .then(response => response.json())
    .then(data => {
      const productContainer = document.querySelector('.container');
      data.forEach(product => {
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
            <h3>${product.Title}</h3>
            <p class="cat">${product.Category}</p>
            <h5>INR ${product.Price}</h5>
            <button class="addtocart" type="submit"  onclick="addToCart(${product.id})">Add to Bag</button> 
        `;
        card.appendChild(prodImg);
        card.appendChild(prodDiv);
        productContainer.appendChild(card);
      })
    })
    .catch(error => {
      console.error('Error fetching data: ', error);
    });
}
fetchproducts();
function fetchcat() {
  fetch('/')
    .then(response => response.json())
    .then(data => {
      let categories = document.querySelector('.categories');
      const uniqueCategories = new Set();
      data.forEach(product => {
        uniqueCategories.add(product.Category);
      });
      Array.from(uniqueCategories).forEach(category => {
        categories.innerHTML += `
              <div class="cat">
                <img src ="./img/${category}.jpg" alt="category">
                <p>${category}</p>
              </div>
              `;

      });

    });
}
fetchcat();
//Emplementing add to cart operation
function addToCart(productId) {
  try {
    fetch('http://127.0.0.1:3000/add-to-cart', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ productId })
    })
      .then(response => {
        if (response.ok) {
          alert('Product added to cart');
        } else {
          alert('Failed to add product to cart.');
        }
      })
      .catch(error => {
        console.error('Error:', error);
      });
  } catch (error) {
    console.error('Error:', error);
  }

}
// cookies handling
let cookies = document.cookie;
let cookie = cookies.split("=");
if (cookies !== "") {
  let login = document.getElementById('login');
  let register = document.getElementById('register');
  login.style.display = 'none';
  register.style.display = 'none';
}
else {
  let logout = document.getElementById('logout');
  let cart = document.getElementById('cart');
  logout.style.display = 'none';
  cart.style.display = 'none';
  dashboard.style.display = 'none';
}
//scroll to top function for UI
window.addEventListener('scroll', function () {
  var scrollToTopBtn = document.getElementById('scrollToTopBtn');
  if (window.scrollY > 200) {
    scrollToTopBtn.classList.add('show');
  } else {
    scrollToTopBtn.classList.remove('show');
  }
});

function scrollToTop() {
  window.scrollTo({ top: 0, behavior: 'smooth' });
}

document.getElementById('scrollToTopBtn').addEventListener('click', scrollToTop);
