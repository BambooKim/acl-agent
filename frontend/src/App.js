import { useState, useEffect } from 'react';
import axios from 'axios';

import Table from 'react-bootstrap/Table';

import logo from './logo.svg';
import './App.css';

const App = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('http://133.186.246.19:8080/api/acl');
        setData(response.data);
      } catch (error) {
        console.error('Error fetching data: ', error)
      }
    }

    fetchData();
  }, []);

  console.log(data)

  return (
    <div>
      <h1>Data Table</h1>
        <Table bordered>
        <thead>
          <tr>
            <th>Id</th>
            <th>Name</th>
            <th>Action</th>
            <th>Direction</th>
            <th>SourceCidr</th>
            <th>SourcePortStart</th>
            <th>SourcePortStop</th>
            <th>DestCidr</th>
            <th>DestPortStart</th>
            <th>DestPortStop</th>
            <th>Protocol</th>
            <th>CreatedAt</th>
            <th>ModifiedAt</th>
          </tr>
        </thead>
        <tbody>
          {data.map((item) => {
            return (
              <tr key={item.id}>
                <td>{item.id}</td>
                <td>{item.name}</td>
                <td>{item.action}</td>
                <td>{item.direction}</td>
                <td>{item.sourceCidr}</td>
                <td>{item.sourcePortStart}</td>
                <td>{item.sourcePortStop}</td>
                <td>{item.destCidr}</td>
                <td>{item.destPortStart}</td>
                <td>{item.destPortStop}</td>
                <td>{item.protocol}</td>
                <td>{item.createdAt}</td>
                <td>{item.modifiedAt}</td>
              </tr>
            )
          })}
        </tbody>
      </Table>
    </div>
  );
}

export default App;
