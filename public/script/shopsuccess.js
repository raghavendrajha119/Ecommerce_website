function seedetails(productId){
    window.location.href = `/product.html?id=${productId}`;
}
function fetchitems() {
    fetch('/get-bought-products') 
        .then(response => response.json())
        .then(data => {
            const Container = document.getElementById('Container');
            Container.innerHTML = '';
            data.forEach(item => {
                const Item = document.createElement('div');
                Item.classList.add('Item');
                const details = document.createElement('div');
                details.classList.add('details');
                details.innerHTML = `
                <h3 class="Title">${item.Title}</h3>
                <p class="Price">Price: INR ${item.TotalAmount}</p>
                <div class="Quantiy_buttons">
                    <p class="Quantity">Quantity: ${item.Quantity}</p>
                <div class="buttons">
                    <button class="Details" type="submit" onclick="seedetails('${item.ID}')">Details</button>
                </div>
                `;
                const Productimg = document.createElement('div');
                Productimg.classList.add('img');
                Productimg.innerHTML = `
                    <img src="./img/${item.Title}.jpg" alt="${item.Title}">
                `;
                Item.appendChild(Productimg);
                Item.appendChild(details);
                Container.appendChild(Item);
            });
        })
        .catch(error => {
            console.error('Error fetching bought products: ', error);
        });
}
fetchitems();
