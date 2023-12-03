import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

import { useState, useEffect } from 'react';
import axios from 'axios';

const AclForm = () => {
  const [ formData, setFormData ] = useState({
    name: '',
    protocol: 'icmp',
    action: 'deny',
    direction: 'ingress',
    sourceCidr: '',
    sourcePortStart: '',
    sourcePortStop: '',
    destCidr: '',
    destPortStart: '',
    destPortStop: ''
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    })
  }

  const handleDropdownChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const requestData = {
        name: formData.name,
        protocol: formData.protocol,
        action: formData.action,
        direction: formData.direction,
        sourceCidr: formData.sourceCidr,
        sourcePortStart: Number(formData.sourcePortStart),
        sourcePortStop: Number(formData.sourcePortStop),
        destCidr: formData.destCidr,
        destPortStart: Number(formData.destPortStart),
        destPortStop: Number(formData.destPortStop)
      };

      const apiUrl = 'http://133.186.246.19:8080/api/acl'

      const response = await axios.post(apiUrl, requestData);

      console.log('Server response: ', response.data)

      setFormData({
        name: '',
        protocol: '',
        action: '',
        direction: '',
        sourceCidr: '',
        sourcePortStart: 0,
        sourcePortStop: 0,
        destCidr: '',
        destPortStart: 0,
        destPortStop: 0
      });
    } catch (error) {
      console.error('Error sending POST request: ', error);
    }
  }
  
  return (
    <Form onSubmit={handleSubmit}>
      <Form.Group className='mb-3' controlId='aclFormName'>
        <Form.Label>Name</Form.Label>
        <Form.Control name='name' value={formData.name} onChange={handleChange} type="text" placeholder="Acl Name" />
      </Form.Group>
      <Form.Group className='mb-3' controlId='aclFormProtocol'>
        <Form.Label>Protocol</Form.Label>
        <Form.Select name='protocol' value={formData.protocol} defaultValue='icmp' onChange={handleDropdownChange}>
          <option value="icmp">icmp</option>
          <option value="tcp">tcp</option>
          <option value="udp">udp</option>
        </Form.Select>
      </Form.Group>
      <Form.Group className='mb-3' controlId='aclFormAction'>
        <Form.Label>Action</Form.Label>
        <Form.Select name='action' value={formData.action} defaultValue='deny' onChange={handleDropdownChange}>
          <option value="permit">permit</option>
          <option value="deny">deny</option>
          <option value="reflect">reflect</option>
        </Form.Select>
      </Form.Group>
      <Form.Group className='mb-3' controlId='aclFormDirection'>
        <Form.Label>Direction</Form.Label>
        <Form.Select name='direction' value={formData.direction} defaultValue='ingress' onChange={handleDropdownChange}>
          <option value="ingress">ingress</option>
          <option value="egress">egress</option>
        </Form.Select>
      </Form.Group>
      <Form.Group className='mb-3' controlId='aclFormSourceCidr'>
        <Form.Label>Source Network</Form.Label>
        <Form.Control name='sourceCidr' value={formData.sourceCidr} onChange={handleChange} type="text" placeholder="0.0.0.0/0" />
      </Form.Group>
      <Form.Group as={Row} className='mb-3' controlId='aclFormSourcePort'>
        <Form.Label>Source Port Range</Form.Label>
        <Col>
          <Form.Control name='sourcePortStart' value={formData.sourcePortStart} onChange={handleChange} type='text' placeholder='Port Range Start'/>
        </Col>
        <Col>
          <Form.Control name='sourcePortStop' value={formData.sourcePortStop} onChange={handleChange} type='text' placeholder='Port Range Stop'/>
        </Col>
      </Form.Group>
      <Form.Group className='mb-3' controlId='aclFormDestCidr'>
        <Form.Label>Destination Network</Form.Label>
        <Form.Control name='destCidr' value={formData.destCidr} onChange={handleChange} type="text" placeholder="0.0.0.0/0" />
      </Form.Group>
      <Form.Group as={Row} className='mb-3' controlId='aclFormDestPort'>
        <Form.Label>Destination Port Range</Form.Label>
        <Col>
          <Form.Control name='destPortStart' value={formData.destPortStart} onChange={handleChange} type='text' placeholder='Port Range Start'/>
        </Col>
        <Col>
          <Form.Control name='destPortStop' value={formData.destPortStop} onChange={handleChange} type='text' placeholder='Port Range Stop'/>
        </Col>
      </Form.Group>
      <div className='d-grid gap-1'>
        <Button type='submit'>
          Create
        </Button>
      </div>
    </Form>
  );
}

export default AclForm;