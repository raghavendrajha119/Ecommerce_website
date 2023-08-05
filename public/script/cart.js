// Assuming you have a function to handle the 'Add to Cart' button click event
function addToCart(productId) {
  fetch('add-to-cart', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ productId: productId }),
  })
    .then((response) => response.json())
    .then((data) => {
      // Handle the response if needed
      console.log(data);
    })
    .catch((error) => {
      // Handle errors if any
      console.error('Error:', error);
    });
}
