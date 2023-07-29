document.addEventListener('DOMContentLoaded', function () {
  let products = document.querySelector('.products');
  async function fetchProducts(url) {
    try {
      let data = await fetch(url);
      let response = await data.json();
      for (let i = 0; i < response.length; i++) {
        let description = response[i].description;
        products.innerHTML += `
            <div class="card">
                <h3 class="Title">${response[i].title}</h3>
                <h4 class="category">${response[i].category}</h4>
                <img src="${response[i].image}" alt="image" class="image">
                <p class="description">${description.length > 20 ? description.substring(0, 80).concat('...more') : description}</p>
                <h5 class="price">$${response[i].price}</h5>
                <button class="cart" onclick="addToCart(${response[i].id})">Add to cart</button>
            </div>
        `;
      }
    } catch (error) {
      console.error('Error:', error);
    }
  }
  fetchProducts('https://fakestoreapi.com/products');
});
document.addEventListener('DOMContentLoaded', function () {
  let categories = document.querySelector('.categories');
  async function fetchcategories(url) {
    let data = await fetch(url);
    let response = await data.json();
    for (let i = 0; i < response.length; i++) {
      let description = response[i].description;
      categories.innerHTML += `
          <div class="one_cat">
            <img src="${response[i].image}" alt="cat">
            <p>${response[i].name}</p>
          </div>
      `;
    }
  };
  fetchcategories('https://api.escuelajs.co/api/v1/categories');
});
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
// emplementing search operation
let search = document.getElementById('search');
search.addEventListener("keyup", (event) => {
  const { value } = event.target;
  const searchQuery = value.toLowerCase();
  let products = document.querySelector('.products');
  products.innerHTML = '';
  async function fetchProducts(url) {
    try {
      let data = await fetch(url);
      let response = await data.json();
      const ids = [];
      for (let i = 0; i < response.length; i++) {
        let description = response[i].description;
        let title = response[i].title;
        let categories = response[i].category;
        let titlecheck = title.toLowerCase();
        let desccheck = description.toLowerCase();
        let catcheck = categories.toLowerCase();
        if ((titlecheck.includes(searchQuery)) || (desccheck.includes(searchQuery)) || (catcheck.includes(searchQuery))) {

          if (ids.includes(response[i].id)) {

          }
          else {
            products.innerHTML += `
            <div class="card">
                <h3 class="Title">${response[i].title}</h3>
                <h4 class="category">${response[i].category}</h4>
                <img src="${response[i].image}" alt="image" class="image">
                <p class="description">${description.length > 20 ? description.substring(0, 80).concat('...more') : description}</p>
                <h5 class="price">$${response[i].price}</h5>
                <button class="cart" onclick="addToCart(${response[i].id})">Add to cart</button>
            </div>
        `;
          }
          ids.push(response[i].id);
        }
      }
    } catch (error) {
      console.error('Error:', error);
    }
  }
  fetchProducts('https://fakestoreapi.com/products');
});
function makeallcolorless(){
  let categories = document.getElementsByClassName('Cat');
for (const category of categories){
  category.style.color = "white";
}
}
// emplementing categories
let categories = document.getElementsByClassName('Cat');
for (const category of categories){
  category.addEventListener('click',()=>{
    makeallcolorless();
    category.style.color = "#ff7846";
    let products = document.querySelector('.products');
    products.innerHTML = '';
    async function fetchProducts(url) {
      try {
        let data = await fetch(url);
        let response = await data.json();
        for (let i = 0; i < response.length; i++) {
          let description = response[i].description;
          let catname = category.textContent.toLowerCase();
          if (response[i].category.includes(catname)){
            products.innerHTML += `
            <div class="card">
                <h3 class="Title">${response[i].title}</h3>
                <h4 class="category">${response[i].category}</h4>
                <img src="${response[i].image}" alt="image" class="image">
                <p class="description">${description.length > 20 ? description.substring(0, 80).concat('...more') : description}</p>
                <h5 class="price">$${response[i].price}</h5>
                <button class="cart" onclick="addToCart(${response[i].id})">Add to cart</button>
            </div>
        `;
          }
        }
      }catch (error) {
        console.error('Error:', error);
      } 
      }
      fetchProducts('https://fakestoreapi.com/products');
  })

}
