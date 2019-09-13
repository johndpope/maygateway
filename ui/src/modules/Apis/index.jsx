import React, { useEffect, useState } from 'react'
import Styled from './styled'
import ApiProvider from 'core/providers/apis'

const Apis = () => {
  const [apis, setApis] = useState([])

  const getApis = () => {
    ApiProvider.getAll()
      .then(res => setApis(res))
  }

  useEffect(() => {
    getApis()
  }, [])

  return (
    <Styled.Wrapper>
      {apis && apis.map(({ ID, name }) => <span key={ID}>{name}</span>)}
    </Styled.Wrapper>
  )
}

export default Apis