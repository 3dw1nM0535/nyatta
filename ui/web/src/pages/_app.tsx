import type { AppProps } from 'next/app'
import Head from 'next/head'
import { UserProvider } from '@auth0/nextjs-auth0/client'
import { getCookie } from 'cookies-next'
import { ChakraProvider } from '@chakra-ui/react'

import { theme } from '@styles'

import { SearchListingProvider } from '../lib/views/listings/providers/search-listings'
import { OnboardingProvider } from '../lib/views/landlord/providers/property-onboarding'

import { ApolloProvider, ApolloClient } from '@apollo/client'
import { AuthProvider } from '@auth'
import { createClient } from '../lib/apollo/createClient'
import Layout from '@layout'

export default function App({ Component, pageProps }: AppProps) {
  const jwt = getCookie('jwt')
  const client = createClient(jwt)

  return (
    <UserProvider>
      <AuthProvider>
        <ApolloProvider client={client as ApolloClient<any>}>
          <ChakraProvider theme={theme} cssVarsRoot="body">
            <Head>
              <meta
                name="viewport"
                content="minimum-scale=1, initial-scale=1, width=device-width, shrink-to-fit=no, viewport-fit=cover"
              />
            </Head>
            <SearchListingProvider>
              <OnboardingProvider>
                <Layout>
                  <Component {...pageProps} />
                </Layout>
              </OnboardingProvider>
            </SearchListingProvider>
          </ChakraProvider>
        </ApolloProvider>
      </AuthProvider>
    </UserProvider>
  )
}
