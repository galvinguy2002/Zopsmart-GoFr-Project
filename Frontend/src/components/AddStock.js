import React, { useState } from 'react';

const AddStock = () => {
  const [stock, setStock] = useState({ symbol: '', price: '', volume: '' });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setStock({ ...stock, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('stockapi/add', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(stock),
      });
      if (response.ok) {
        console.log('Stock added successfully!');
        
      } else {
        console.error('Failed to add stock');
        
      }
    } catch (error) {
      console.error('Error:', error);
      
    }
  };

  return (
    <div>
      <h2>Add Stock</h2>
      <form onSubmit={handleSubmit}>
        <input type="text" name="symbol" value={stock.symbol} onChange={handleChange} placeholder="Symbol" />
        <input type="text" name="price" value={stock.price} onChange={handleChange} placeholder="Price" />
        <input type="text" name="volume" value={stock.volume} onChange={handleChange} placeholder="Volume" />
        <button type="submit">Add Stock</button>
      </form>
    </div>
  );
};

export default AddStock;
