import React, { useState, useEffect } from 'react';

const ViewStocks = () => {
  const [stocks, setStocks] = useState([]);

  useEffect(() => {
    const fetchStocks = async () => {
      try {
        const response = await fetch('stockapi/view');
        if (response.ok) {
          const data = await response.json();
          setStocks(data);
        } else {
          console.error('Failed to fetch stocks');
          
        }
      } catch (error) {
        console.error('Error:', error);
        
      }
    };
    fetchStocks();
  }, []);

  return (
    <div>
      <h2>Stocks List</h2>
      <ul>
        {stocks.map((stock) => (
          <li key={stock.id}>
            {stock.symbol} - Price: {stock.price} - Volume: {stock.volume}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ViewStocks;
