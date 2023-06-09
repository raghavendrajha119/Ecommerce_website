const API_URL = 'https://fakestoreapi.com/products';
    
    // Function to fetch and display products
    async function fetchProducts() {
      try {
        const response = await fetch(API_URL);
        const products = await response.json();
        
        const productContainer = document.getElementById('product-list');
        productContainer.innerHTML = '';
        
        products.forEach(product => {
          const card = createProductCard(product);
          productContainer.appendChild(card);
        });
      } catch (error) {
        console.error('Error fetching products:', error);
      }
    }
    
    // Function to create a product card
    function createProductCard(product) {
      const card = document.createElement('div');
      card.classList.add('product');
      
      const title = document.createElement('h2');
      title.textContent = product.title;
      
      const description = document.createElement('p');
      description.textContent = product.description;
      
      const price = document.createElement('p');
      price.textContent = `$${product.price}`;
      
      const editButton = document.createElement('button');
      editButton.textContent = 'Edit';
      editButton.classList.add('button');
      editButton.addEventListener('click', () => showEditForm(product));
      
      const deleteButton = document.createElement('button');
      deleteButton.textContent = 'Delete';
      deleteButton.classList.add('button');
      deleteButton.addEventListener('click', () => deleteProduct(product.id));
      
      card.appendChild(title);
      card.appendChild(description);
      card.appendChild(price);
      card.appendChild(editButton);
      card.appendChild(deleteButton);
      
      return card;
    }
    
    // Function to add a new product
    async function addProduct(name, description, price) {
      try {
        const response = await fetch(API_URL, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            title: name,
            description: description,
            price: parseFloat(price),
          }),
        });
        
        if (response.ok) {
            const product = await response.json();
            const card = createProductCard(product);
            const productContainer = document.getElementById('product-list');
            productContainer.appendChild(card);
            clearForm(); // Clear the form inputs
        } else {
          console.error('Error adding product:', response.statusText);
        }
      } catch (error) {
        console.error('Error adding product:', error);
      }
    }
    
    // Function to delete a product
    async function deleteProduct(productId) {
      try {
        const response = await fetch(`${API_URL}/${productId}`, {
          method: 'DELETE',
        });
        
        if (response.ok) {
          fetchProducts(); // Refresh the product list
        } else {
          console.error('Error deleting product:', response.statusText);
        }
      } catch (error) {
        console.error('Error deleting product:', error);
      }
    }
    // Function to show the edit form
    function showEditForm(product) {
        isEditing = true;
        currentProductId = product.id;
        
        document.getElementById('edit-id').value = product.id;
        document.getElementById('edit-name').value = product.title;
        document.getElementById('edit-description').value = product.description;
        document.getElementById('edit-price').value = product.price;
        
        document.getElementById('edit-form').style.display = 'block';
        document.getElementById('product-list').style.display = 'none';
    }
    // Function to update a product
    async function updateProduct(id, name, description, price) {
        try {
          const response = await fetch(`${API_URL}/${id}`, {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              title: name,
              description: description,
              price: parseFloat(price),
            }),
          });
          
          if (response.ok) {
            fetchProducts(); // Refresh the product list
            cancelEdit(); // Hide the edit form
          } else {
            console.error('Error updating product:', response.statusText);
          }
        } catch (error) {
          console.error('Error updating product:', error);
        }
    }
    // Function to cancel editing and hide the edit form
    function cancelEdit() {
        isEditing = false;
        currentProductId = null;
        
        document.getElementById('edit-form').style.display = 'none';
        document.getElementById('product-list').style.display = 'block';
      }
      
      // Handle form submission for adding or updating a product
      document.getElementById('product-form').addEventListener('submit', (event) => {
        event.preventDefault();
        
        const name = document.getElementById('name-input').value;
        const description = document.getElementById('description-input').value;
        const price = document.getElementById('price-input').value;
        
        if (isEditing) {
          updateProduct(currentProductId, name, description, price);
        } else {
          addProduct(name, description, price);
        }
    });
    // Handle cancel button click for edit form
    document.getElementById('cancel-edit').addEventListener('click', () => {
        cancelEdit();
    });
    
    // Fetch and display the initial list of products
    fetchProducts();