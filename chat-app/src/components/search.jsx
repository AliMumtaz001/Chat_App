import React, { useState } from 'react';
// import { getUsers } from '../apis/apislist/api';
import { getUsers } from '../apis/apislist/api';

const SearchUser = ({ onUserSelect }) => {
  const [searchQuery, setSearchQuery] = useState('');
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(false);
  const token = localStorage.getItem('accessToken');

  const handleSearch = async (e) => {
    const query = e.target.value;
    setSearchQuery(query);
    if (query.length > 2) {
      setLoading(true);
      try {
        const result = await getUsers(query, token);
        console.log('API Response:', result);
        if (result.data) {
          setUsers(result.data.users || []);
        } else {
          setUsers([]);
          console.error('Search failed:', result.error);
        }
      } catch (error) {
        console.error('Error fetching users:', error);
        setUsers([]);
      }
      setLoading(false);
    } else {
      setUsers([]);
    }
  };

  const handleSelect = (user) => {
    onUserSelect(user);
    setSearchQuery('');
    setUsers([]);
  };

  return (
    <div className="p-4">
      <input
        type="text"
        value={searchQuery}
        onChange={handleSearch}
        placeholder="Search users..."
        className="w-full p-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      {loading && <p className="text-gray-600">Loading...</p>}
      <ul className="mt-2">
        {users.map((user) => (
          <li
            key={user.id}
            onClick={() => handleSelect(user)}
            className="p-2 hover:bg-gray-200 cursor-pointer"
          >
            {user.username}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default SearchUser;