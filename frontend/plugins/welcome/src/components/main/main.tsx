import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../listappoint';
import Button from '@material-ui/core/Button';


 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
const profile = { givenName: 'รายการข้อมูลนัดตรวจพิเศษ' };
const HeaderCustom = {
    minHeight: '50px',
  };
  
 return (
    
    <Page theme={pageTheme.home} >
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
     <Header
       style={HeaderCustom}
       title={`${profile.givenName}`}
       subtitle=""
     >    
     
        <div>
        <Link component={RouterLink} to="/login">
          <Button variant="contained" color="primary">
             ออกจากระบบ
         </Button>
         </Link>
          <br></br>
          <Link component={RouterLink} to="/addappoint">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              เพิ่มข้อมูลนัดตรวจพิเศษ
            </Button>
          </Link>
         </div>
    </Header>
     <Content>
       <ContentHeader title="ข้อมูลการนัดตรวจพิเศษ">
       </ContentHeader>
       <div>
       </div>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;