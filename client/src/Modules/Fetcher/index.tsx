import { Form, Formik, FormikErrors, FormikHelpers, FormikValues } from 'formik'
import Api from 'Modules/Api'
import Input from 'Modules/Form/Input'
import Submit from 'Modules/Form/Submit'
import React, { useState } from 'react'
import ReactJson from 'react-json-view'
import { FormGroup, Label } from 'reactstrap'
import 'Styles/Fetcher.scss'

interface Props {}

interface SubmitType {
  value: FormikValues
  helpers: FormikHelpers<FormikValues>
}

const Fetcher: React.FC<Props> = () => {
  const [rawData, setRawData] = useState({})
  const initialValue: FormikValues = {
    url: ''
  }

  const onSubmit = async (
    values: FormikValues,
    helpers: FormikHelpers<FormikValues>
  ) => {
    try {
      const { data } = await Api().post('/fetcher', values)
      setRawData(data)
    } catch (e) {
      helpers.setStatus({
        error: 'terjadi kesalahan pada server'
      })
    }
  }

  const onValidate = (values: FormikValues) => {
    let err: FormikErrors<FormikValues> = {}
    if (values.url === '') {
      err.url = 'Harap isi URL Berita'
    }
    return err
  }

  return (
    <Formik
      initialValues={initialValue}
      validate={onValidate}
      onSubmit={onSubmit}
    >
      <div className="fetcher">
        <Form>
          <FormGroup>
            <Label>URL Berita</Label>
            <Input
              name="url"
              type="text"
              placeholder="http://detik.com/news/..."
            />
          </FormGroup>
          <FormGroup>
            <Submit color="primary">Dapatkan Berita</Submit>
          </FormGroup>
        </Form>

        <h3>Result</h3>
        <div className="result-data">
          <ReactJson
            src={rawData}
            theme="monokai"
            enableClipboard={false}
            displayDataTypes={false}
            displayObjectSize={false}
          />
        </div>
      </div>
    </Formik>
  )
}

export default Fetcher
