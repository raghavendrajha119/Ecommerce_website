document.addEventListener('DOMContentLoaded', function(){
    let products = document.querySelector('.products'); 
    async function fetchProducts(url){
        let data = await fetch (url);
        let response = await data.json();
        for (let i = 0; i< response.length; i++){
            let description = response[i].description;
            products.innerHTML += `
            <div class="card">
                <h3 class="Title">${response[i].title}</h3>
                <h4 class="category">${response[i].category}</h4>
                <img src="${response[i].image}" alt="image" class="image">
                <p class="description">${description.length > 20 ? description.substring(0,80).concat('...more') : description}</p>
                <h5 class="price">$${response[i].price}</h5>
                <button class="cart">Add to cart</button>
            </div>
        `;
        }  
    };
    fetchProducts('https://fakestoreapi.com/products');
});
document.addEventListener('DOMContentLoaded', function(){
  let categories = document.querySelector('.categories'); 
  async function fetchcategories(url){
      let data = await fetch (url);
      let response = await data.json();
      for (let i = 0; i< response.length; i++){
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
window.addEventListener('scroll', function() {
    var scrollToTopBtn = document.getElementById('scrollToTopBtn');
    if (window.pageYOffset > 200) {
      scrollToTopBtn.classList.add('show');
    } else {
      scrollToTopBtn.classList.remove('show');
    }
  });
  
  function scrollToTop() {
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }
  
  document.getElementById('scrollToTopBtn').addEventListener('click', scrollToTop);