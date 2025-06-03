import React, { useState } from 'react';
import SearchUser from './search';
// import MapRecentUsers from './users';
import ChatSection from './chatbox';

const Main = () => {
  const [selectedUser, setSelectedUser] = useState(null);

  const handleUserSelect = (user) => {
    setSelectedUser(user);
  };

  return (
    <div className="flex h-screen">
      <div className="w-1/4 border-r">
        <SearchUser onUserSelect={handleUserSelect} />
        {/* <MapRecentUsers onUserSelect={handleUserSelect} /> */}
      </div>
      <div className="w-3/4">
        <ChatSection selectedUser={selectedUser} />
      </div>
    </div>
  );
};

export default Main;