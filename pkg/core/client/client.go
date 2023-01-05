package client

import (
    "bytes"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

func Post(url string, body []byte) error {
    client := &http.Client{}
    client.Timeout = time.Second * 5

    //Header部分
    header := http.Header{}
    header.Set("Content-Length", "10000")
    header.Add("Content-Type", "application/json")

    //リクエストの作成
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        return err
    }

    req.Header = header

    resp, err := client.Do(req)
    if err != nil {
        log.Println(err)
        return err
    }
    defer resp.Body.Close()

    _, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func Get(url string, body []byte) error {
    client := &http.Client{}
    client.Timeout = time.Second * 5

    //Header部分
    header := http.Header{}
    header.Set("Content-Length", "10000")
    header.Add("Content-Type", "application/json")

    //リクエストの作成
    req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(body)))
    if err != nil {
        return err
    }

    req.Header = header

    resp, err := client.Do(req)
    if err != nil {
        log.Println(err)
        return err
    }
    defer resp.Body.Close()

    _, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func Put(url string, body []byte) error {
    client := &http.Client{}
    client.Timeout = time.Second * 5

    //Header部分
    header := http.Header{}
    header.Set("Content-Length", "10000")
    header.Add("Content-Type", "application/json")

    //リクエストの作成
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(body)))
    if err != nil {
        return err
    }

    req.Header = header

    resp, err := client.Do(req)
    if err != nil {
        log.Println(err)
        return err
    }
    defer resp.Body.Close()

    _, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func Delete(url string, body []byte) error {
    client := &http.Client{}
    client.Timeout = time.Second * 5

    //Header部分
    header := http.Header{}
    header.Set("Content-Length", "10000")
    header.Add("Content-Type", "application/json")

    //リクエストの作成
    req, err := http.NewRequest("DELETE", url, bytes.NewBuffer([]byte(body)))
    if err != nil {
        return err
    }

    req.Header = header

    resp, err := client.Do(req)
    if err != nil {
        log.Println(err)
        return err
    }
    defer resp.Body.Close()

    _, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}
