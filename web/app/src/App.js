import React, { useState } from 'react';
import "bootstrap/dist/css/bootstrap.css";
import { Container, Row, Col, Form, Button, InputGroup, Card } from 'react-bootstrap';



const App = () => {
  const [orderId, setOrderId] = useState('');
  const [responseJson, setResponseJson] = useState('{\n}');

  const handleFetchOrder = async () => {
    try {
      const response = await fetch(`http://localhost:8080/order/${orderId}`);
      const json = await response.json();
      setResponseJson(JSON.stringify(json, null, 4));
    } catch (error) {
      console.error('Error fetching order:', error);
      setResponseJson("Error fetching order")
    }
  };

  return (
    <div className="App bg-light min-vh-100">
      <Container className="py-4">
      <Row>
        <Col md={8} className="mx-auto mt-4">
          <h1 className="text-center mb-3">Order Fetcher</h1>
          <Card>
            <Card.Body>
              <Form>
                <InputGroup className="mb-3">
                  <Form.Control
                    placeholder="Enter Order ID"
                    value={orderId}
                    onChange={(e) => setOrderId(e.target.value)}
                  />
                  <Button variant="primary" onClick={handleFetchOrder}>
                    Fetch
                  </Button>
                </InputGroup>
              </Form>
            </Card.Body>
          </Card>
        </Col>
      </Row>
      <Row>
        <Col md={8} className="mx-auto">
          <Card className="mt-3">
            <Card.Body>
              <Card.Title>Response JSON:</Card.Title>
              <pre>{responseJson}</pre>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
    </div>
  );
};

export default App;

