import React from 'react'
import { Container, Navbar, NavbarBrand } from 'reactstrap'
interface Props {}

const PageContainer: React.FC<Props> = props => {
  return (
    <div className="app">
      <Navbar color="dark" dark>
        <Container className="app-container">
          <NavbarBrand href="#">Yo-Crawler</NavbarBrand>
        </Container>
      </Navbar>
      {props.children}
    </div>
  )
}

export default PageContainer
