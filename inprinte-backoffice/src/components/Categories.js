import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, Create,} from 'react-admin';
  
  export const CategoriesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="name" />
      </Datagrid>
    </List>
  );
  
  export const CategoriesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="id" disabled />
        <TextInput source="name" />
      </SimpleForm>
    </Edit>
  );

  export const CategoriesCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="name" />
      </SimpleForm>
    </Create>
  );