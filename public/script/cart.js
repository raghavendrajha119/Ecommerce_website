function seedetails(productId){
    window.location.href = `/product.html?id=${productId}`;
}
function proceedtobuy() {
    fetch('/checkout', {
        method: 'POST', // Assuming this is the endpoint for checkout
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => response.json())
    .then(data => {
        console.log("Checkout response:", data);
        if (data.message === "Purchase completed successfully.") {
            const url = `/shopsuccess.html`;
            window.location.href = url;
        } else {
            console.log("Checkout failed:", data.message);
        }
    })
    .catch(error => {
        console.error('Error during checkout: ', error);
    });
}

function removeprod(productId){
    fetch('/remove-from-cart',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            productId: parseInt(productId)
        })
    })
    .then(response => response.json())
    .then(data => {
        alert(data.message);
        fetchitems();
    })
    .catch(error => {
        alert('Error removing product: ' + error);
    });
}
function updateQuantityInDatabase(productId,newQuantity) {
    fetch('/update-cart-quantity', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            productId: parseInt(productId),
            quantity: parseInt(newQuantity)
        })
    })
    .then(response => response.json())
    .then(data => {
        if (data.message === "Out of stock") {
            alert(`Out of stock! Available: ${data.available}, Requested: ${data.requested}`);
        } else {
            alert(data.message);
            fetchitems();
        }
    })
    .catch(error => {
        alert('Error updating quantity: ', error);
    });
}

function fetchitems() {
    fetch('/get-cart')
        .then(response => response.json())
        .then(data => {
            console.log(data)
            const cartContainer = document.getElementById('cartContainer');
            cartContainer.innerHTML = '';
            const cartproceedbuy = document.createElement('div');
            cartproceedbuy.classList.add('proceedcart');
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
                        <input type="number" class="quantity-input" value="${item.Quantity}" min="1">
                    </div>
                    </p>
                    <button class="update-quantity" data-product-id="${item.ID}">Update Quantity</button>
                </div>
                <div class="buttons">
                    <button class="Details" type="submit" onclick="seedetails('${item.ID}')">Details</button>
                    <button class="Buynow">Buy Now</button>
                    <button class="removefromcart" id="removeprod" onclick="removeprod('${item.ID}')">Remove</button>
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
                const updateQuantityBtn = cartItem.querySelector('.update-quantity');
                updateQuantityBtn.addEventListener('click', () => {
                    const productId = updateQuantityBtn.getAttribute('data-product-id');
                    const newQuantity = cartItem.querySelector('.quantity-input').value;
                    updateQuantityInDatabase(productId, newQuantity);
                });
            });
            cartproceedbuy.innerHTML=`
                <button class="cartbuy" type="submit" onclick="proceedtobuy('')">Proceed to checkout</button>
                `;
            cartContainer.appendChild(cartproceedbuy);
        })
        .catch(error => {
            console.error('Error fetching cart items: ', error);
        });
}
fetchitems();
