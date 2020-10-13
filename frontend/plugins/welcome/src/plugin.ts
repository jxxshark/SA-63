import { createPlugin} from '@backstage/core';
import welcomepage from './components/addappoint';
import createuser from './components/listappoint';
import main from './components/main';
import login from './components/login';
export const plugin = createPlugin({
  id: 'appoint',
  register({ router }) {
    router.registerRoute('/addappoint', welcomepage);
    router.registerRoute('/p', createuser);
    router.registerRoute('/main',main);
    router.registerRoute('/login',login);
  },
  
});
