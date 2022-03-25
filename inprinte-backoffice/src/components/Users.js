import {List, Datagrid, TextField, EmailField, UrlField, Edit, SimpleForm, TextInput, Create,} from 'react-admin';
  
  export const UserList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="firstname" />
        <TextField source="lastname" />
        <TextField source="role" />
        <TextField source="password" />
        <EmailField source="email" />
        <TextField source="phone" />
        <TextField source="is_alive" />
        <TextField source="address.street" label="Street" />
        <TextField source="address.city" label="City" />
        <TextField source="address.state" label="State" />
        <TextField source="address.country" label="Country" />
        <TextField source="address.zipCode" label="ZipCode" />
      </Datagrid>
    </List>
  );
  
  export const UserEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="id" disabled />
        <TextInput source="firstname" />
        <TextInput source="lastname" />
        <TextInput source="email" />
        <TextInput source="phone" />
        <TextInput source="is_alive" />
        <TextInput source="role.name" label="role"/>
        <TextInput source="address.street" label="Street" />
        <TextInput source="address.city" label="City" />
        <TextInput source="address.state" label="State" />
        <TextInput source="address.country" label="Country" />
        <TextInput source="address.zipCode" label="ZipCode" />
      </SimpleForm>
    </Edit>
  );

  export const UserCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="firstname" />
        <TextInput source="lastname" />
        <TextInput source="email" />
        <TextInput source="phone" />
        <TextInput source="password" />
        <TextInput source="is_alive" />
        <TextInput source="role.name" label="role"/>
      </SimpleForm>
    </Create>
  );