import { useState, useEffect } from 'react';
import axios from 'axios';

import Table from 'react-bootstrap/Table';

import AclForm from './AclForm';

import 'bootstrap/dist/css/bootstrap.css';


const convertDate = (dateStr) => {
  const rawDate = new Date(dateStr)

  const year = rawDate.getFullYear();
  const month = (rawDate.getMonth() + 1).toString().padStart(2, '0');
  const day = rawDate.getDate().toString().padStart(2, '0');
  const hours = rawDate.getHours().toString().padStart(2, '0');
  const minutes = rawDate.getMinutes().toString().padStart(2, '0');

  return `${year}년 ${month}월 ${day}일 ${hours}시 ${minutes}분`
}


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
      <h1>ACL Agent</h1>
      <Table bordered>
        <thead>
          <tr>
            <th>Name</th>
            <th>Protocol</th>
            <th>Action</th>
            <th>Direction</th>
            <th>SourceCidr</th>
            <th>SourcePortStart</th>
            <th>SourcePortStop</th>
            <th>DestCidr</th>
            <th>DestPortStart</th>
            <th>DestPortStop</th>
            
            {/* <th>CreatedAt</th>
            <th>ModifiedAt</th> */}
          </tr>
        </thead>
        <tbody>
          {data.map((item) => {
            return (
              <tr key={item.id}>
                <td>{item.name}</td>
                <td>{item.protocol}</td>
                <td>{item.action}</td>
                <td>{item.direction}</td>
                <td>{item.sourceCidr}</td>
                <td>{item.sourcePortStart}</td>
                <td>{item.sourcePortStop}</td>
                <td>{item.destCidr}</td>
                <td>{item.destPortStart}</td>
                <td>{item.destPortStop}</td>
                
                {/* <td>{convertDate(item.createdAt)}</td>
                <td>{convertDate(item.modifiedAt)}</td> */}
              </tr>
            )
          })}
        </tbody>
      </Table>
      <AclForm/>
    </div>
  );
}

export default App;
