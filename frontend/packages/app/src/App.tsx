import React, { FC } from 'react';
import {
  createApp,
  OAuthRequestDialog,
} from '@backstage/core';
import { apis } from './apis';
import * as plugins from './plugins';
import { Routes, Navigate } from 'react-router';


const app = createApp({
  apis,
  plugins: Object.values(plugins),
});

const AppProvider = app.getProvider();
const AppRouter = app.getRouter();
const deprecatedAppRoutes = app.getRoutes();




const App: FC<{}> = () => (
  <AppProvider>
    
    <OAuthRequestDialog />
    <AppRouter>
     
        
        <Routes>
          <Navigate key="/" to="/login"/>
         
          
          
          {deprecatedAppRoutes}
        </Routes>
     
    </AppRouter>
  </AppProvider>
);

export default App;
