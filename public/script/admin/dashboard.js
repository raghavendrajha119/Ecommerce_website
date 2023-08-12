const statsContainer = document.querySelector('.stats');

fetch('/admin/dashboard', {
  method: 'GET',
})
  .then(response => response.json())
  .then(data => {
    const customerStats = data.totalUsers;
    const adminUserStats = data.totalAdminUsers;
    const productStats = data.totalProducts;

    const customerStatElement = document.createElement('div');
    customerStatElement.className = 'customer';
    customerStatElement.innerHTML = `<p>Present Users: ${customerStats}</p>`;
    statsContainer.appendChild(customerStatElement);
    const adminUserStatElement = document.createElement('div');
    adminUserStatElement.className = 'admin-users';
    adminUserStatElement.innerHTML = `<p>Admin Users: ${adminUserStats}</p>`;
    statsContainer.appendChild(adminUserStatElement);
    const productStatElement = document.createElement('div');
    productStatElement.className = 'product';
    productStatElement.innerHTML = `<p>Present Products: ${productStats}</p>`;
    statsContainer.appendChild(productStatElement);
  })
  .catch(error => {
    console.error('Error fetching data:', error);
  });
