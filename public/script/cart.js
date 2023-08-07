function seedetails(productId){
    window.location.href = `/product.html?id=${productId}`;
}
function fetchitems() {
    fetch('/get-cart')
        .then(response => response.json())
        .then(data => {
            console.log(data)
            const cartContainer = document.getElementById('cartContainer');
            cartContainer.innerHTML = '';
            data.forEach(item => {
                const cartItem = document.createElement('div');
                cartItem.classList.add('cart-item');
                const Cartdetails = document.createElement('div');
                Cartdetails.classList.add('cartdetails');
                Cartdetails.innerHTML = `
            <h3 class="Title">${item.Title}</h3>
            <p class="cat">${item.Category}</p>
            <p class="Price">Price: INR ${item.Price}</p>
            <div class="Quantiy_buttons">
            <p class="Quantity">Quantity: ${item.Quantity}
            <div class="quantity-dial">
                <input type="number" class="quantity-input" value="1" min="1">
            </div>
            </div>
            <div class="buttons">
            <button class="Details" type="submit" onclick="seedetails('${item.ID}')">Details</button>
                <button class="Buynow">Buy Now</button>
                <button class="removefromcart" id="removeprod">Remove</button>
            </div>
            `;
                const Productimg = document.createElement('div');
                Productimg.classList.add('img');
                Productimg.innerHTML = `
                <img src="./img/${item.Title}.jpg" alt="${item.Title}">
                `;
                cartItem.appendChild(Productimg);
                cartItem.appendChild(Cartdetails);
                cartContainer.appendChild(cartItem);
            });
        })
        .catch(error => {
            console.error('Error fetching cart items: ', error);
        });
}
fetchitems();
