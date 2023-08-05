function fetchData() {
    fetch('/products')
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
            <h2>${product.Title}<h2>
            <p class="cat">${product.Category}</p>
            <h5>Price: INR ${product.Price}</h5>
            <p class="desc">${product.Description}</p>
            <button class="addtocart" type="submit">Add to Bag</button> 
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
fetchData()
//carousel operations
// let slidecontainer = document.querySelector('.carousel-slide');
// let prevbutton = document.querySelector('.carousel-prev');
// let nextButton = document.querySelector('.carousel-next');
// let images = document.querySelectorAll('.carousel-slide img');
// let counter = 0;
// function prevSlide() {
//     counter--;
//     if (counter < 0) {
//         counter = images.length - 1;
//     }
//     updateSlide();
// }
// function nextSlide() {
//     counter++;
//     if (counter >= images.length) {
//         counter = 0;
//     }
//     updateSlide();
// }
// function updateSlide() {
//     slidecontainer.style.transform = `translateX(-${counter * 100}%)`;
// }

// prevbutton.addEventListener('click', prevSlide);
// nextbutton.addEventListener('click', nextSlide);