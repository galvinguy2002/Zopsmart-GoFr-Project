import React, { useState } from 'react';

const UpdateStock = () => {
  const [stock, setStock] = useState({ id: '', symbol: '', price: '', volume: '' });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setStock({ ...stock, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`stockapi/update/${stock.id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(stock),
      });
      if (response.ok) {
        console.log('Stock updated successfully!');
        
      } else {
        console.error('Failed to update stock');
        
      }
    } catch (error) {
      console.error('Error:', error);
      
    }
  };

  return (
    <div>
      <h2>Update Stock</h2>
      <form onSubmit={handleSubmit}>
        <input type="text" name="id" value={stock.id} onChange={handleChange} placeholder="Stock ID" />
        <input type="text" name="symbol" value={stock.symbol} onChange={handleChange} placeholder="Symbol" />
        <input type="text" name="price" value={stock.price} onChange={handleChange} placeholder="Price" />
        <input type="text" name="volume" value={stock.volume} onChange={handleChange} placeholder="Volume" />
        <button type="submit">Update Stock</button>
      </form>
    </div>
  );
};

export default UpdateStock;
