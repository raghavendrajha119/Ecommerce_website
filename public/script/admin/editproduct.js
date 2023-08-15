const editForm = document.getElementById('editForm');
const urlParams = new URLSearchParams(window.location.search);
const productID = urlParams.get('id');
fetch(`/admin/edit-product?id=${productID}`)
    .then(response => response.json())
    .then(product => {
        document.getElementById('title').value = product.Title;
        document.getElementById('price').value = product.Price;
        document.getElementById('description').value = product.Description;
        document.getElementById('category').value = product.Category;
        document.getElementById('quantity').value = product.Quantity;
    })
    .catch(error => {
        console.error('Error fetching product data: ', error);
    });

editForm.addEventListener('submit', async event => {
    event.preventDefault();
    const formData = new FormData(editForm);
    try {
        const response = await fetch(`/admin/update-product/${productID}`, {
            method: 'POST',
            body: formData
        });
        if (response.ok) {
            alert("Product updated successfully!");
        } else {
            alert("Failed to add product");
        }
    } catch (error) {
        console.error("Error updating product:", error);
    }
});