import React, { useState } from 'react';

const DeleteStock = () => {
  const [stockId, setStockId] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`stockapi/d/${stockId}`, {
        method: 'GET',
      });
      if (response.ok) {
        console.log('Stock deleted successfully!');
        
      } else {
        console.error('Failed to delete stock');
        
      }
    } catch (error) {
      console.error('Error:', error);
      
    }
  };

  return (
    <div>
      <h2>Delete Stock</h2>
      <form onSubmit={handleSubmit}>
        <input type="text" value={stockId} onChange={(e) => setStockId(e.target.value)} placeholder="Stock ID" />
        <button type="submit">Delete Stock</button>
      </form>
    </div>
  );
};

export default DeleteStock;
