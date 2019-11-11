import Fetcher from 'Modules/Fetcher'
import PageContainer from 'Modules/PageContainer'
import React from 'react'
import { Container } from 'reactstrap'
import 'Styles/App.scss'

const App: React.FC = () => {
  return (
    <PageContainer>
      <Container className="app-container">
        <Fetcher />
      </Container>
    </PageContainer>
  )
}

export default App
