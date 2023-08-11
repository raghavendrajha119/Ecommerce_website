document.addEventListener("DOMContentLoaded", function () {
    const form = document.querySelector("form");

    form.addEventListener("submit", async function (event) {
        event.preventDefault();

        const formData = new FormData(form);
        try {
            const response = await fetch("/admin/add-products", {
                method: "POST",
                body: formData,
            });

            if (response.ok) {
                alert("Product added successfully!");
                form.reset();
            } else {
                alert("Failed to add product");
            }
        } catch (error) {
            console.error("Error adding product:", error);
        }
    });
});
