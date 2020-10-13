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
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import { EntPatient } from '../../api/models/EntPatient';
import { EntSpecializeddiag } from '../../api/models/EntSpecializeddiag';
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

/*const initialUserState = {
 name: 'System Analysis and Design',
 age: 20,
};
*/

export default function Create() {
 const classes = useStyles();
 const profile = { givenName: 'ระบบนัดตรวจพิเศษ' };
 const api = new DefaultApi();
 //const [user, setUser] = useState(initialUserState);
 const [users, setUsers] = useState<EntUser[]>([]);
 const [patients, setPatients] = useState<EntPatient[]>([]);
 const [specials, setSpecials] = useState<EntSpecializeddiag[]>([]);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 const [loading, setLoading] = useState(true);


 const [userid, setUserid] = useState(Number);
 const [patientid, setPatientid] = useState(Number);
 const [specialid, setSpecialid] = useState(Number);
 const [datetime, setDatetime] = useState(String);

 useEffect(() => {

   const getPatients = async () => {

     const pa = await api.listPatient({ limit: 10, offset: 0 });
     setLoading(false);
     setPatients(pa);
   };
   getPatients();

   const getUsers = async () => {

   const us = await api.listUser({ limit: 10, offset: 0 });
     setLoading(false);
     setUsers(us);
   };
   getUsers();

   const getSpecials = async () => {

    const sp = await api.listSpecializeddiag({ limit: 10, offset: 0 });
      setLoading(false);
      setSpecials(sp);
    };
    getSpecials();

 }, [loading]);

    const handleDatetimeChange = (event: any) => {
        setDatetime(event.target.value as string);
    };

 const CreateSpecializedappoint = async () => {
     const specializedappoint = {
        date              : datetime,
        patientID         : patientid,
        specializeddiagID : specialid,
        userID            : userid,
     }
   const res:any = await api.createSpecializedappoint({ specializedappoint : specializedappoint});
   setStatus(true);
   if (res.id != ''){
     setAlert(true);
   } else {
     setAlert(false);
   }

   const timer = setTimeout(() => {
     setStatus(false);
   }, 1000);
 };

 const patient_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setPatientid(event.target.value as number);
   };

  const user_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUserid(event.target.value as number);
   };

  const special_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSpecialid(event.target.value as number);
   };

 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`${profile.givenName || 'to Backstage'}`}
       subtitle="บันทึกข้อมูลนัดตรวจพิเศษ."
     ></Header>
     <Content>
       <ContentHeader title="Create">
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 This is a success alert — check it out!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>
       <div className={classes.root}>
         <form noValidate autoComplete="off">
<table>
<tr><td align="right">Name</td><td>
           <FormControl
             fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <InputLabel id="name_id-label">Name_ID</InputLabel>
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
{/*
           <FormControl
             fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="age"
               label="Age"
               variant="outlined"
               type="number"
               size="medium"
               value={user.age}
               onChange={handleInputChange}
             />
           </FormControl>
 */}

<tr><td align="right">Patient</td><td>

           <FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="patient_id-label">Patient_ID</InputLabel>
            <Select
              labelId="patient_id-label"
              label="Patient"
              id="patient_id"
              value={patientid}
              onChange={patient_id_handleChange}
              style = {{width: 600}}
            >
              {patients.map((item:EntPatient)=>
              <MenuItem value={item.id}>{item.name}</MenuItem>)}
            </Select>
 </FormControl>
</td></tr>
<tr><td align="right">Special</td><td>
 <div>
<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            
            <InputLabel id="special_id-label">Special_ID</InputLabel>
            <Select
              labelId="special_id-label"
              label="Special"
              id="special_id"
              value={specialid}
              onChange={special_id_handleChange}
              style = {{width: 600}}
            >
              {specials.map((item:EntSpecializeddiag)=>
              <MenuItem value={item.id}>{item.specializeddiacnostictype}</MenuItem>)}
            </Select>
 </FormControl>
 </div>
 </td></tr>
 <tr><td align="right">DateTime</td><td>
 <div>
            <FormControl className={classes.margin} >
              <TextField
                id="date"
                label="DateTime"
                type="date"
                value={datetime}
                onChange={handleDatetimeChange}
    //defaultValue="2017-05-24"
                className={classes.textField}
                InputLabelProps={{
                shrink: true,
                }}
              />
            </FormControl>

 </div>
</td></tr>
 </table>

           <div className={classes.margin}>
             <Button
               onClick={() => {
                 CreateSpecializedappoint();
                 
               }}
               component={RouterLink}
               to="/main"
               variant="contained"
               
               color="primary"
             >
               บันทึกข้อมูล
             </Button>
             <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/p"
               variant="contained"
             >
               Back
             </Button>
           </div>
         </form>
       </div>
     </Content>
   </Page>
 );
}