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


export default function Create() {
 const classes = useStyles();
 const profile = { givenName: 'ระบบนัดตรวจพิเศษ' };
 const api = new DefaultApi();
 const [users, setUsers] = useState<EntUser[]>([]);
 const [patients, setPatients] = useState<EntPatient[]>([]);
 const [specials, setSpecials] = useState<EntSpecializeddiag[]>([]);
 const [load, setLoading] = useState(true);


 const [userid, setUserid] = useState(Number);
 const [patientid, setPatientid] = useState(Number);
 const [specialid, setSpecialid] = useState(Number);
const [date, setDatetime] = useState(String);


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

 }, [load]);

    const handleDatetimeChange = (event: any) => {
        setDatetime(event.target.value as string);
    };
   
 const CreateSpecializedappoint = async () => {
     const specializedappoint = {
        date              : date + "T00:00:00+07:00",
        patientID         : patientid,
        specializeddiagID : specialid,
        userID            : userid,
     }
    await api.createSpecializedappoint({ specializedappoint : specializedappoint});
 

   
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
   <Page theme={pageTheme.service}>
     <Header
       title={`${profile.givenName || 'to Backstage'}`}
       subtitle={<span style={{fontSize:'550px',color: 'whitesmoke',font:'message-box'}}>{"บันทึกข้อมูลตรวจพิเศษ"}</span>}
     ></Header>
     <Content>
       <ContentHeader title="สำหรับการบันทึกข้อมูลนัดตรวจพิเศษ">
       </ContentHeader>
       <div className={classes.root}>
         <form noValidate autoComplete="off">
<table>
<tr><td align="right">แพทย์mที่ทำการนัด</td><td>
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

<tr><td align="right">ผู้ป่วย</td><td>

           <FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="patient_id-label">ชื่อผู้ป่วย</InputLabel>
            <Select
              labelId="patient_id-label"
              label="Patient"
              id="patient_id"
              value={patientid}
              onChange={patient_id_handleChange}
              style = {{width: 600}}
            >
              {patients.map((item:EntPatient)=>
              <MenuItem value={item.id}>{item.patientname}</MenuItem>)}
            </Select>
 </FormControl>
</td></tr>
<tr><td align="right">รายการนัดตรวจพิเศษ</td><td>
 <div>
<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            
            <InputLabel id="special_id-label">รายการนัดตรวจพิเศษ</InputLabel>
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
 
 <tr><td align="right">วันที่ทำการบันทึก</td><td>
 <div>
            <FormControl className={classes.margin} >
              <TextField
                id="date"
                label="วันที่"
                type="date"
                value={date}
                onChange={handleDatetimeChange}
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