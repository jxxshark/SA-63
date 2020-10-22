import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { DefaultApi } from '../../api/apis';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import { EntUser } from '../../api/models/EntUser';
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
     margin: theme.spacing(3),
   },
   withoutLabel: {
     marginTop: theme.spacing(3),
   },
   textField: {
     width: '25ch',
   },
 }),
);


export default function Create() {
 const classes = useStyles();
 const profile = { givenName: 'ระบบนัดตรวจพิเศษ' };
 const api = new DefaultApi();
 const [users, setUsers] = useState<EntUser[]>([]);
 const [loading, setLoading] = useState(true);


 const [userid, setUserid] = useState(Number);



 useEffect(() => {

   const getUsers = async () => {

   const us = await api.listUser({ limit: 10, offset: 0 });
     setLoading(false);
     setUsers(us);
   };
   getUsers();

  
 }, [loading]);

   
 

  const user_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUserid(event.target.value as number);
   };

  
   
 return (
   <Page theme={pageTheme.service}>
     <Header
       title={`${profile.givenName || 'to Backstage'}`}
       subtitle={<span style={{fontSize:'550px',color: 'whitesmoke',font:'message-box'}}>{"บันทึกข้อมูลตรวจพิเศษ"}</span>}
     ></Header>
     <Content>
       <ContentHeader title="เข้าสู่ระบบ">
       </ContentHeader>
       <div className={classes.root}>
         <form noValidate autoComplete="off">
<table>
<tr><td align="right">แพทย์ที่ทำการนัด</td><td>
           <FormControl
             fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <InputLabel id="name_id-label">ชื่อแพทย์ที่ทำการนัด</InputLabel>
             <Select
               labelId="user_id-label"
               label="User"
               id="user_id"
               value={userid}
               onChange={user_id_handleChange}
               style = {{width: 600}}
               >
               {users.map((item:EntUser)=>
               <MenuItem value={item.id}>{item.username}</MenuItem>)}
             </Select>
           </FormControl>
           </td></tr>
 </table>
           
           <div className={classes.margin}>
           <center>
             <Button
               onClick={() => {
                 CreateSpecializedappoint();
                 
               }}
               component={RouterLink}
               to="/main"  
               
               variant="contained"
               
               color="secondary"
             >
               บันทึก
             </Button>
             </center>
             
           </div>
         </form>
       </div>
     </Content>
   </Page>
 );
}