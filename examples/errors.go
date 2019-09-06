func CopyFile(src, dst string) error {
    fin, err := os.Open(src)
    if err != nil {
        return err
    }
    defer fin.Close()
    fout, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|...)
    if err != nil {
        return err
    }
    defer fout.Close()
    _, err = io.Copy(fout, fin)
    return err
}

func CopyFile(src, dst string) error {
    fin := check os.Open(src)
    defer fin.Close()
    fout := check os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|...)
    defer fout.Close()
    check io.Copy(fout, fin)
    return nil
}

func CopyFile(src, dst string) error {
    fin := try(os.Open(src))
    defer fin.Close()
    fout := try(os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644))
    defer fout.Close()
    _, err = io.Copy(fout, fin)
    return err
}

func CopyFile(src, dst string) error {
    fin, err := os.Open(src)
    if err != nil {
        return fmt.Errorf(“Cannot open source file: %w”, err)
    }
    defer fin.Close()
    fout, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf(“Cannot open destination file: %w”, err)
    }
    defer fout.Close()
    If _, err = io.Copy(fout, fin); err != nil {
        return fmt.Errorf(“Cannot copy content from %v to %v: %w”, src, dst, err)
    }
    return nil
}

func As(err error, target interface{}) bool {...}
func Is(err, target error) bool {...}
func Unwrap(err error) error {...}
