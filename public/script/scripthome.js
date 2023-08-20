function seedetails(productId){
  window.location.href = `/product.html?id=${productId}`;
}
function addToCart(productId){
  fetch('/add-to-cart',{
    method: 'POST',
    headers:{
      'Content-Type':'application/json'
    },
    body: JSON.stringify({ productId:productId })
  })
  .then(response => {
    if (!response.ok){
      throw new Error('Failed to add item to cart.');
    }
  })
  .catch (error => {
    console.error('Error adding item to cart:',error);
  });
}
function fetchproducts() {
  fetch('/home')
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
        console.log(product.ID);
        prodDiv.innerHTML = `
            <h3>${product.Title}</h3>
            <p class="cat">${product.Category}</p>
            <h5>INR ${product.Price}</h5>
            <button class="seedetails" type="submit" onclick="seedetails('${product.ID}')">Details</button>
            <div class="buttons">
              <button class="addtocart" type="submit"  onclick="addToCart(${product.ID})">Add to Bag</button>
              <button class="buynow" type="submit">Buy Now</button>
            </div>
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
  fetch('/home')
      .then(response => response.json())
      .then(data => {
          let categories = document.querySelector('.categories');
          const uniqueCategories = new Set();
          data.forEach(product => {
              uniqueCategories.add(product.Category);
          });
          console.log(uniqueCategories)
          Array.from(uniqueCategories).forEach(category => {
              categories.innerHTML += `
                  <div class="cat">
                      <button class="category-btn" data-category="${category}">
                          <img src="./img/${category}.jpg" alt="category">
                          <p>${category}</p>
                      </button>
                  </div>
              `;
          });

          // Attach event listener to category buttons
          document.querySelectorAll('.category-btn').forEach(button => {
              button.addEventListener('click', () => {
                  const selectedCategory = button.getAttribute('data-category');
                  filterProductsByCategory(selectedCategory);
              });
          });
      });
}

function filterProductsByCategory(category) {
  fetch('/home')
      .then(response => response.json())
      .then(data => {
          const productContainer = document.querySelector('.container');
          productContainer.innerHTML = '';

          data.forEach(product => {
              if (product.Category === category) {
                const card = document.createElement('div');
                card.classList.add('card');
                const prodImg = document.createElement('div');
                prodImg.classList.add('img');
                prodImg.innerHTML = `
                <img src="./img/${product.Title}.jpg" alt="${product.Title}">
                `;
                const prodDiv = document.createElement('div');
                prodDiv.classList.add('details');
                console.log(product.ID);
                prodDiv.innerHTML = `
                    <h3>${product.Title}</h3>
                    <p class="cat">${product.Category}</p>
                    <h5>INR ${product.Price}</h5>
                    <button class="seedetails" type="submit" onclick="seedetails('${product.ID}')">Details</button>
                    <div class="buttons">
                      <button class="addtocart" type="submit"  onclick="addToCart(${product.ID})">Add to Bag</button>
                      <button class="buynow" type="submit">Buy Now</button>
                    </div>
                `;
                card.appendChild(prodImg);
                card.appendChild(prodDiv);
                productContainer.appendChild(card);
              }
          });
      })
      .catch(error => {
          console.error('Error fetching products: ', error);
      });
}
document.addEventListener('DOMContentLoaded', () => {
  fetchcat();

  document.querySelectorAll('.category-btn').forEach(button => {
      button.addEventListener('click', () => {
          const selectedCategory = button.textContent;
          filterProductsByCategory(selectedCategory);
      });
  });
});
//implementing search operation
document.getElementById('search-form').addEventListener('submit',function(event){
  event.preventDefault();
  const searchInput = document.getElementById('search').value;
  searchProducts(searchInput)
})
function searchProducts(query) {
  fetch('/home')
      .then(response => response.json())
      .then(data => {
          const productContainer = document.querySelector('.container');
          productContainer.innerHTML = '';

          data.forEach(product => {
              if (product.Category.toLowerCase().includes(query.toLowerCase()) ||
                  product.Title.toLowerCase().includes(query.toLowerCase())) {

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
                      <button class="seedetails" type="submit" onclick="seedetails('${product.ID}')">Details</button>
                      <div class="buttons">
                          <button class="addtocart" type="submit" onclick="addToCart(${product.ID})">Add to Bag</button>
                          <button class="buynow" type="submit">Buy Now</button>
                      </div>
                  `;
                  card.appendChild(prodImg);
                  card.appendChild(prodDiv);
                  productContainer.appendChild(card);
              }
          });
      })
      .catch(error => {
          console.error('Error searching products: ', error);
      });
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
  let orders =document.getElementById('Orders')
  logout.style.display = 'none';
  cart.style.display = 'none';
  dashboard.style.display = 'none';
  orders.style.display='none'
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
